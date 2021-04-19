import javax.swing.*;
import java.awt.*;
import java.util.ArrayList;

public class CanvasPanel extends JPanel {
    public ArrayList<Integer> x1 = new ArrayList<>();
    public ArrayList<Integer> x2 = new ArrayList<>();
    public ArrayList<Integer> y1 = new ArrayList<>();
    public ArrayList<Integer> y2 = new ArrayList<>();
    public ArrayList<Integer> fuckYou = new ArrayList<>();

    protected void paintComponent(Graphics g) {
        super.paintComponent(g);

        for (int i = 0; i < PictureForm.length; i++) {
            if (fuckYou.get(i) == 0) {
                g.setColor(Color.BLUE);
            } else if (fuckYou.get(i) == 1) {
                g.setColor(Color.RED);
            } else {
                g.setColor(Color.YELLOW);
            }
            g.drawLine(x1.get(i), y1.get(i), x2.get(i), y2.get(i));
        }
    }
}
