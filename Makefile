Day7:
	echo "Starting"
	cd 2023 && go test . -v -run TestDay7 
Day11:
	cd 2023 && go test . -v -run TestDay11 
Test:
	cd util && go test . -v -run TestShortestPath
