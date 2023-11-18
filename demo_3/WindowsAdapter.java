public class WindowsAdapter extends ROG{

    private ROG rog;

    public WindowsAdapter(ROG rog) {
        this.rog = rog;
    }

    public void insertInSquarePort(){
        System.out.println("Converting {Square} to {Circle} port");
        rog.insertInCirclePort();
   }
}
