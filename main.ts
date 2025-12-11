/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2025-12-10
 * @fileoverview This program is a simple calculator.
 */

console.log("Welcome to my calculator program.");

console.log("Which operation would you like to perform today?");
console.log("a. add");
console.log("b. subtract");
console.log("c. multiply");
console.log("d. divide");
console.log("e. absolute value");
console.log("f. round");
console.log("g. exponent");
console.log("h. square root");

let choice = prompt("What operation do you want to choose: ");

if (choice == "a") {
  let x = Number(prompt("Enter first number: "));
  let y = Number(prompt("Enter second number: "));
  console.log(x + " + " + y + " = " + (x + y));
}

else if (choice == "b") {
  let x = Number(prompt("Enter first number: "));
  let y = Number(prompt("Enter second number: "));
  console.log(x + " - " + y + " = " + (x - y));
}

else if (choice == "c") {
  let x = Number(prompt("Enter first number: "));
  let y = Number(prompt("Enter second number: "));
  console.log(x + " * " + y + " = " + (x * y));
}

else if (choice == "d") {
  let x = Number(prompt("Enter numerator: "));
  let y = Number(prompt("Enter denominator: "));
  console.log(x + " / " + y + " = " + (x / y));
}

else if (choice == "e") {
  let x = Number(prompt("Enter number: "));
  console.log("Absolute value of " + x + " is " + Math.abs(x));
}

else if (choice == "f") {
  let x = Number(prompt("Enter number: "));
  console.log("Rounded value is " + Math.round(x));
}

else if (choice == "g") {
  let a = Number(prompt("Enter base: "));
  let b = Number(prompt("Enter exponent: "));
  console.log(a + " raised to " + b + " = " + Math.pow(a, b));
}

else if (choice == "h") {
  console.log("In order to calculate the square root, you will need a non-negative value:");
  let x = Number(prompt("Enter value: "));
  console.log("The square root of " + x + " is " + Math.sqrt(x));
}

else {
  console.log("Invalid choice.");
}
