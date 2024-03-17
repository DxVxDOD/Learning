import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Scanner;

public class Reading {
    public static void main(String[] args) {

        ArrayList<String> lines = new ArrayList<>();

        Scanner scannerPick = new Scanner(System.in);
        String pick = scannerPick.nextLine();

        try(Scanner scanner = new Scanner(Paths.get(pick))) {
            while (scanner.hasNextLine()){
                String row = scanner.nextLine();
                lines.add(row);
                System.out.println(row);
            }
        }catch (Exception e){
            System.out.println("Error: " + e.getMessage());
        }

        try(Scanner scanner = new Scanner(Paths.get("Hello.txt"))) {
            while (scanner.hasNextLine()){
                String row = scanner.nextLine();
                lines.add(row);
                System.out.println(row);
            }
        }catch (Exception e){
            System.out.println("Error: " + e.getMessage());
        }

        System.out.println(lines.size() + " lines");

    }

}
