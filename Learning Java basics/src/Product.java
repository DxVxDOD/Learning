public class Product {
    private String name;
    private double price;
    private int quantity;
    public Product(String initialName, double initialPrice, int initialQuantity){
        this.price = initialPrice ;
        this.name = initialName ;
        this.quantity = initialQuantity ;
    }

    public void printProduct() {
        System.out.println(this.name + "Price: " + this.price + ", " + this.quantity + "pcs");
    }

    public void increasePrice(double amount){
        this.price += amount;
    }

    public void resetPrice(){
        this.price = 0;
    }

    @Override
    public String toString(){
        return this.name + "Hahahah";
    }
}
