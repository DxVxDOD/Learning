import java.util.Scanner;

public class Example {
  public static void main(String[] args) {

    Scanner scanner = new Scanner(System.in);

    System.out.println("Yes or no?");

    String message = scanner.nextLine();

    if (message.equals("yes")) {
      System.out.println("HAahahhaha");
    } else if (message.equals("no")) {
      System.out.println("sdlakndlasnldsaknd");
    } else {
      System.out.println("Respcet the rules");
    }

    scanner.close();
  }
}
