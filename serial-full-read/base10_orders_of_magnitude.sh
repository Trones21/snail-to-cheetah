 #!/bin/bash
 
 echo "Cleanup.... this may take awhile if dir is large (>1M files)"
rm -rf ~/filegen_out && mkdir ~/filegen_out

 
 echo "Create Files"
 go run ../filegen/main.go -dir ~/filegen_out/10k -max 5000 -min 100 -num 10000
 go run ../filegen/main.go -dir ~/filegen_out/100k -max 5000 -min 100 -num 100000
 go run ../filegen/main.go -dir ~/filegen_out/1M -max 5000 -min 100 -num 1000000

echo "Run Benchmarks"
go run ./main.go -path ~/filegen_out/10k
go run ./main.go -path ~/filegen_out/100k
go run ./main.go -path ~/filegen_out/1M