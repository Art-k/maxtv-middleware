open("/home/art-k/PROJECT/MAXTV-TOOLS/database/list_of_files.sh", "w").write("")
open("/home/art-k/PROJECT/MAXTV-TOOLS/database/import_dumps.sh", "w").write("")

infile = open("/home/art-k/PROJECT/MAXTV-TOOLS/database/list_of_files.txt", "r")
for line in infile:
    line = line.replace("\n", "")
    print(line)
    tmp = "scp -P 885 euadmin@maxtvmedia.com:/srv/backup2/sql/maxtv_cms_live/2021/03/21/" + \
          line + " ./backup/" + line + "\n" + "gzip -d ./backup/" + line + "\n"
    open("/home/art-k/PROJECT/MAXTV-TOOLS/database/list_of_files.sh", "a").write(tmp)

    tmp1 = "/usr/bin/mysql -pmax1029tv --database=maxtv_live --user=art --host=0.0.0.0 --port=6603" + \
           " < /home/art-k/PROJECT/MAXTV-TOOLS/database/backup/" + line[:-3] + "\n"
    open("/home/art-k/PROJECT/MAXTV-TOOLS/database/import_dumps.sh", "a").write(tmp1)
