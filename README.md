# SeqSearch
引物探针序列占比分析工具

# 概述
基于KMP和Hamming算法计算引物或探针序列在reads中的占比

# 使用

- 克隆仓库

```shell
git clone https://github.com/51cat/SeqSearch.git
cd SeqSearch
./psearch -h
```
- 基本命令
```shell
./psearch \
    -input ./test_data/test.fastq \
    -format fastq \
    -target_fa ./test_data/target.fa \
    -out ./test_data/test_out.tsv \
    -mismatch 2
```

# 详细使用步骤

```shell
./psearch -h

```

# 测试脚本

```shell
./test.sh
```
