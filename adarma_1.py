arr = [
	[0, 0, 0],
	[0, 0, 0],
	[0, 0, 0],
];

while True:
	inp = input("Give input: ");
	coord = inp.split(",")

	turn = True
	
	if turn:
		arr[coord[0]][coord[1]] = 1;
		turn = False;
	else:
		arr[coord[0]][coord[1]] = 2;
		turn = True
	print(arr)
