git add -A
git commit -m $1
git push https://github.com/alax-mx/mtproto.git

# git add -A              #添加所有修改
# git commit -m "提交修改" #提交修改
# git push url main       #同步到main分支到git服务器
# git pull url main       #拉去git服务器上的main分支到本地
# git branch              #列出本地所有分支
# git branch test         #创建test分支
# git checkout test       #切换本地分支为 test分支
# git branch -b test      #创建test分支，并切换到test分支
# git branch -d test      #删除test分支
# git clone -b brhA url   #指定brhA分支
# git reset --soft HEAD^  回退到上一个版本
# git reset --soft HEAD^^ 回退到上上个版本