function bubble_sort(arr){
  let sorted = false
  let temp = 0
  let un_i = arr.length - 1
  while(!sorted){
    sorted = true
    for(let i = 0; i < un_i; i++){
      if(arr[i] > arr[i + 1]){
          sorted = false;
          temp = arr[i]
          arr[i] = arr[i + 1]
          arr[i + 1] = temp
      }
    }
    un_i = un_i - 1
  }
  return arr
}

console.log(bubble_sort([65, 55, 45, 69, 35, 25, 15, 10]));

function hasDuplicateValue(array) {
var steps = 0;
for(var i = 0; i < array.length; i++) {
for(var j = 0; j < array.length; j++) {
steps++;
if(i !== j && array[i] == array[j]) {
return true;
}
}
}
console.log(steps);
return false;
}

hasDuplicateValue([65, 55, 45, 69, 35, 25, 15, 10])

