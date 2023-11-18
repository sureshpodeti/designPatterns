public class demo {

    public static void main(String args[]){
        Mac mac = new Mac();
        mac.insertInSquarePort();

        ROG rog = new ROG();

        WindowsAdapter windowsAdapter = new WindowsAdapter(rog);
        windowsAdapter.insertInSquarePort();
    }
}
