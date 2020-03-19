# word2vec
1.this is word2vec with golang 
2.直接运行creatmodel.go，会对已经含有的语料库进行训练生成模型，运行命令： go run creatmodel.go 
3.在原有go版本的word2vec基础上，添加了对词向量计算的工具word-calc
使用说明可以参见下面这个例子： 以word-calc为例 进入cmd文件夹下面运行 ./word-calc -model modelpath(模型路径) -add 词语 -v，
即可进行查看每个词语转换而成的词向量
