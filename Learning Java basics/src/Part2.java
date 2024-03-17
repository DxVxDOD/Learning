import java.util.ArrayList;

public class Part2 {
    public static void main(String[] args) {

        ArrayList<String> textArray = new ArrayList<>();

        textArray.add("Hello");
        textArray.add("!");
        textArray.add("World");
        textArray.add("!");

        textArray.remove(1);

        System.out.println(textArray);
        System.out.println(textArray.get(2));

    }
}
