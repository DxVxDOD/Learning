function sortPoker(john, uncle) {
  let suits_order = [];
  const suits = ["♦", "♠", "♣", "♥"];
  const letter_value = new Map([
    ["A", "14"],
    ["K", "13"],
    ["Q", "12"],
    ["J", "11"],
  ]);
  const number_letter = new Map([
    ["14", "A"],
    ["13", "K"],
    ["12", "Q"],
    ["11", "J"],
  ]);
  let card_map = new Map();
  let res = "";
  let curr_suit = "";

  uncle.split("").forEach((char) => {
    if (suits.includes(char) && !suits_order.includes(char)) {
      suits_order.push(char);
    }
  });

  console.log(suits_order);

  for (let i = 0; i < suits_order.length; i++) {
    card_map.set(suits_order[i], []);

    console.log([...card_map.entries()]);
  }

  for (let char of john.split("")) {
    if (suits_order.includes(char)) {
      curr_suit = char;
    } else {
      const char_int = parseInt(char);
      if (typeof char_int === "number" && !isNaN(char_int)) {
        let arr = card_map.get(curr_suit);
        if (char_int === 1) {
          arr.push("10");
          card_map.set(curr_suit, arr);
        } else if (char_int === 0) {
          continue;
        } else {
          console.log(curr_suit);
          let arr2 = card_map.get(curr_suit);
          console.log(arr2);
          arr2.push(char);
          card_map.set(curr_suit, arr2);
        }
      } else {
        let arr = card_map.get(curr_suit);
        arr.push(letter_value.get(char));
        card_map.set(curr_suit, arr);
      }
    }
  }

  for (let [key, value] of card_map) {
    card_map.set(
      key,
      value.sort((a, b) => parseInt(a) - parseInt(b)),
    );
  }

  for (let [key, value] of card_map) {
    for (let v of value) {
      res += key;
      if (parseInt(v) > 10) {
        res += number_letter.get(v);
      } else {
        res += v;
      }
    }
  }

  return res;
}

console.log(
  sortPoker(
    "♦6♥2♠3♦5♠J♣Q♠K♣7♦2♣5♥5♥10♠A",
    "♠2♠3♠5♥J♥Q♥K♣8♣9♣10♦4♦5♦6♦7",
  ),
);
