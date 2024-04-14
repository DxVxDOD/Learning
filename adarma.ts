const arr = [
	[0, 0, 0],
	[0, 0, 0],
	[0, 0, 0],
];

while (true) {
	const input = prompt("Give input: ");
	const coord = input?.split(",")!;

	let turn = true;

	if (turn) {
		arr[coord[0]][coord[1]] = 1;
		turn = false;
	} else {
		arr[coord[0]][coord[1]] = 2;
		turn = true;
	}

	console.log(arr);
}
