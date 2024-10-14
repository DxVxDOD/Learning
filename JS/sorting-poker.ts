function sortPoker(jhon: string, uncle: string) {
  let suit_order: string[] = [];
  const suits = ["♣", "♠", "♦", "♥"];

  uncle.split("").forEach((char) => {
    if (suits.includes(char) && !suit_order.includes(char)) {
      suit_order.push(char);
    }
  });

  let letter_value = new Map<string, string>([
    ["A", "14"],
    ["K", "13"],
    ["Q", "12"],
    ["J", "11"],
  ]);
  let number_letter = new Map<string, string>([
    ["14", "A"],
    ["13", "K"],
    ["12", "Q"],
    ["11", "J"],
  ]);

  let card_map = new Map<string, string[]>();
  for (let i = 0; i < suit_order.length; i++) {
    card_map.set(suit_order[i], []);
  }

  let res = "";
  let curr_suit = "";

  for (let char of jhon.split("")) {
    if (suit_order.includes(char)) {
      curr_suit = char;
    } else {
      const char_int = parseInt(char);
      if (typeof char_int === "number" && !isNaN(char_int)) {
        let arr = card_map.get(curr_suit);
        if (char_int === 1) {
          arr?.push("10");
          card_map.set(curr_suit, arr!);
        } else if (char_int === 0) {
          continue;
        } else {
          arr?.push(char);
          card_map.set(curr_suit, arr!);
        }
      } else {
        let arr = card_map.get(curr_suit);
        arr?.push(letter_value.get(char)!);
        card_map.set(curr_suit, arr!);
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

  console.log([...card_map.entries()]);
  console.log(res);
  return res.split("");
}

console.log(
  sortPoker(
    "♦6♥2♠3♦5♠J♣Q♠K♣7♦2♣5♥5♥10♠A",
    "♠2♠3♠5♥J♥Q♥K♣8♣9♣10♦4♦5♦6♦7",
  ),
);

function sortPokers(john: string, uncle: string) {
  let suits_order: string[] = [];
  const suits = ["♦", "♠", "♣", "♥"];
  const letter_value = new Map<string, string>([
    ["A", "14"],
    ["K", "13"],
    ["Q", "12"],
    ["J", "11"],
  ]);
  const number_letter = new Map<string, string>([
    ["14", "A"],
    ["13", "K"],
    ["12", "Q"],
    ["11", "J"],
  ]);
  let card_map = new Map<string, string[]>();
  let res = "";
  let curr_suit = "";

  uncle.split("").forEach((char) => {
    if (suits.includes(char) && !suits_order.includes(char)) {
      suits_order.push(char);
    }
  });

  for (let suit in suits_order) {
    card_map.set(suit, []);
  }

  for (let char of john.split("")) {
    if (suits_order.includes(char)) {
      curr_suit = char;
    } else {
      const char_int = parseInt(char);
      if (typeof char_int === "number" && !isNaN(char_int)) {
        let arr = card_map.get(curr_suit);
        if (char_int === 1) {
          arr?.push("10");
          card_map.set(curr_suit, arr!);
        } else if (char_int === 0) {
          continue;
        } else {
          arr?.push(char);
          card_map.set(curr_suit, arr!);
        }
      } else {
        let arr = card_map.get(curr_suit);
        arr?.push(letter_value.get(char)!);
        card_map.set(curr_suit, arr!);
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
