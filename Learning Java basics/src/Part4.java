public class Part4 {
    public static void main(String[] args) {
        Whistle dog = new Whistle("wau");

        Door door = new Door();
        door.knock();

        System.out.println(dog.sound());

        Product apple = new Product("apple", 10.5, 200);
        apple.printProduct();
        apple.increasePrice(10);
        apple.increasePrice(10);
        apple.printProduct();
        apple.resetPrice();
        apple.printProduct();

        Product orange = new Product("orange", 12.6, 300);
        orange.printProduct();
    }
}