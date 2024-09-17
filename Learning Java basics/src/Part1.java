import java.util.Scanner;

public class Part1 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.println("Print stars:");
        int num = Integer.parseInt(scanner.nextLine());
        int num2 = Integer.parseInt(scanner.nextLine());

        printSquare(num);
        printRectangle(num, num2);
        printChristmasTree(num);

    }

    public static void printStars(int num) {
        int i = 0;
        while (i < num) {

            System.out.print("*");

            i++;
        }
    }


    public static void printDash(int num) {

        int i = 0;
        while (i < num) {
            System.out.print("-");
            i++;
        }

    }
    public static void printChristmasTree(int size) {

        for (int i = 1; i <= size; i++) {

            printDash((size - i));
            printStars(i);
            printStars(i - 1);
            printDash((size - i));
            System.out.println(" ");
        }

        printTrunk(size);

    }

    public static void printTrunk(int size) {
        printDash(size - 2);
        printStars(3);
        printDash(size - 2);
        System.out.println(" ");
        printDash(size - 2);
        printStars(3);
        printDash(size - 2);
    }

    public static void printRectangle(int width, int height) {
        for (int i = 0; i < height; i++) {
            printStars(width);
            System.out.println(" ");
        }
    }

    public static void printSquare(int num) {

        int i = 0;

        while (i < num) {

            printStars(num);
            System.out.println(" ");

            i++;
        }

    }
}
