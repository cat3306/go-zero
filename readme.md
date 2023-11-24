## change

`gozero` 的作者不打算支持`gorm`

自己实现`gorm` 风格的`model`

用法

```bash
goctl model mysql ddl --src /path/user.sql --dir /path/model -m gorm
```



```
goctl model mysql ddl -h

Generate mysql model from ddl
...........
  -m, --mode string       chose gorm style
.................
```

