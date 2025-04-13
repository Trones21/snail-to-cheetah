Paths are just kinda random here, obviously adjust depend on where you are running from.

### Filegen
```bash
go run main.go --num 100 --min 102400 --max 512000 --dir ./test_files/<some_test>
```
- min and max are the sizes in bytes
- please create subdirectories for individual tests



### Serial Full Read
```bash
go run main.go -path ../test_files
```