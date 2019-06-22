bee new jasonpsgo
bee api jasonpsgo -conn="jasonpsgo:jasonpsgo@tcp(127.0.0.1:3306)/jasonpsgo"
bee generate docs
bee run watchall true
cd jasonpsgo
bee run -downdoc=true -gendoc=true    生成Wwagger API
chown -R www:www *



生成models：
bee generate model user
bee generate model user -fields="name:string,age:int"


生成controller:
bee generate controller user


生成view:
bee generate view user


生成文档:
bee generate docs