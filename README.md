# monitorDNS
## 监控指定域名的A记录解析变化, 会统计A记录变化的频率,以及最大,最小,平均值, 并收集所有出现过的A记录

## 选项
* -d 选项 // 必须指定
* -i 间隔 // 单位秒
* -p 是否在控制台打印 // 默认不打印

## 示例
> monitorDNS.exe -d www.baidu.com
> 
> 默认会在运行目录生成一个log.txt文件, 输出域名的解析变化
> 
![JGD7JJTXSCS0X75RD{1Z2W4](https://user-images.githubusercontent.com/52809998/164955704-09ce9189-5cd4-4498-b3ac-0bf3ee6116eb.png)
