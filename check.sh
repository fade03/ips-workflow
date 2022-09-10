
if [ ! -f "./ips" ]; then
	cd ./core && go build -o ips && mv ./ips ../ && cd ../ && ./ips
else
	./ips
fi
