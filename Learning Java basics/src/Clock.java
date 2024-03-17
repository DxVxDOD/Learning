public class Clock {

    public static void main(String[] args) {

        boolean cond = true;

        ClockHand hours = new ClockHand(24);
        ClockHand minutes = new ClockHand(60);
        ClockHand seconds = new ClockHand(60);

        while (cond) {
            // 1. Printing the time
            System.out.println(hours + ":" + minutes + ":" + seconds);

            // 2. Advancing the second hand
            seconds.advance();

            // 3. Advancing the other hands when required
            if (seconds.value() == 0) {
                minutes.advance();

                if (minutes.value() == 0) {
                    hours.advance();
                }
            }
            if(hours.value() > 3){
                cond = false;
            }
        }

    }

}
