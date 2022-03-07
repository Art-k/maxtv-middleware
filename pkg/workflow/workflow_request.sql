SELECT t2.name,t1.*,
       t3.id as order_id,
       t3_parent.id as order_id_parent,
       if(t3.id, t3.order_type, t3_parent.order_type) as order_type,
       t3.sale_date as order_sale_date,
       t3_parent.sale_date as order_sale_date_parent,
       t1.ad_type,
       IF(t3.collected = 1,'YES',IF(t3.collected = -1,'NO','')) as collected,
       IF(t3_parent.collected = 1,'YES',IF(t3_parent.collected = -1,'NO','')) as collected_parent,
       UNIX_TIMESTAMP(t3.sale_date) as ut_osd,
       UNIX_TIMESTAMP(t3_parent.sale_date) as ut_osd_parent,
       (
           SELECT count(t4.`id`) FROM maxtv_company_campaign_notes_general t4 WHERE t4.company_id = t1.company_id
       ) as notes,
       unix_timestamp((t1.email_art_request)) as ut_ear,
       unix_timestamp((t1.email_ad_draft)) as ut_ead,
       unix_timestamp((t1.email_psd_ad_draft)) as ut_epad,
       unix_timestamp((t1.designer_1)) as ut_d1,
       unix_timestamp((t1.designer_2)) as ut_d2,
       unix_timestamp((t1.designer_3)) as ut_d3,
       cats.name as category_name,
       t1.type as campaign_type,
       t1.status,
       (SELECT `pay`.`depositedon` FROM `maxtv_company_payments` `pay` WHERE `order_id` IN(t1.order_id, t3_parent.id) ORDER BY `pay`.`depositedon` DESC LIMIT 1) as deposition_date
FROM maxtv_company_campaigns t1
         inner JOIN maxtv_companies t2 on t2.id = t1.company_id AND t2.exclude_report=0
         LEFT JOIN maxtv_company_orders t3 on t3.id  = t1.order_id
         LEFT JOIN maxtv_company_orders t3_parent on t3_parent.id  = (SELECT order_id FROM maxtv_company_campaigns WHERE id = t1.parent_id)
         LEFT JOIN maxtv_company_categories cats on cats.id  = t1.category_id
WHERE (t1.`status` <> 'cancelled' && t1.`title` NOT LIKE '%cancelled%') AND
    (
            !(DATE(t1.start_date) <= date(NOW()) and date(NOW()) <= date(t1.end_date))
            or
            ((SELECT count(t3.`id`) FROM maxtv_company_campaign_media t3 WHERE  t3.campaign_id = t1.id AND t3.`active`=1) < 1 AND t1.ad_type = "S")
            or
            ((SELECT count(t3.`id`) FROM maxtv_company_campaign_banners t3 WHERE  t3.campaign_id = t1.id AND t3.`active`=1) < 1 AND t1.ad_type = "B")
            or
            (t1.type = 'primary' and t1.email_your_ad_is_up = "0000-00-00 00:00:00")
        )
  and
    !(
                DATE(t1.start_date) > Date(NOW())
            and
                ((SELECT count(t3.`id`) FROM maxtv_company_campaign_media t3 WHERE t3.campaign_id = t1.id) > 0
                    or
                 (SELECT count(t3.`id`) FROM maxtv_company_campaign_banners t3 WHERE t3.campaign_id = t1.id) > 0)
            and
                (t1.type = 'primary' and t1.email_your_ad_is_up != "0000-00-00 00:00:00")
        )
  and
    !(
                t1.type != 'primary'
            and
                ((SELECT count(t3.`id`) FROM maxtv_company_campaign_media t3 WHERE t3.campaign_id = t1.id) > 0
                    or
                 (SELECT count(t3.`id`) FROM maxtv_company_campaign_banners t3 WHERE t3.campaign_id = t1.id) > 0)
            and
                DATE(NOW()) > DATE(t1.end_date)
            and
                t1.end_date != "0000-00-00"
        )
  and
    !(
                t1.type = 'primary'
            and
                ((SELECT count(t3.`id`) FROM maxtv_company_campaign_media t3 WHERE  t3.campaign_id = t1.id) > 0
                    or
                 (SELECT count(t3.`id`) FROM maxtv_company_campaign_banners t3 WHERE  t3.campaign_id = t1.id) > 0)
            and
                DATE(NOW()) > DATE(t1.end_date)
            and
                t1.email_your_ad_is_up != "0000-00-00 00:00:00"
        )
ORDER by
    IF(GREATEST(t1.email_art_request,t1.designer_1,t1.designer_2,t1.email_ad_draft,t1.email_psd_ad_draft) = t1.email_art_request, t1.email_art_request,0) desc,
    IF(GREATEST(t1.designer_1,t1.designer_2,t1.designer_3,t1.email_ad_draft,t1.email_psd_ad_draft) = t1.designer_1, t1.designer_1,0) desc,
    IF(GREATEST(t1.designer_2,t1.designer_3,t1.email_ad_draft,t1.email_psd_ad_draft) = t1.designer_2, t1.designer_2,0) desc,
    IF(GREATEST(t1.designer_3,t1.email_ad_draft,t1.email_psd_ad_draft) = t1.designer_2, t1.designer_2,0) desc,
    t1.email_ad_draft desc