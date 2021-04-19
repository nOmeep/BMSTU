import javax.swing.*;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;

public class PictureForm {
    public static int length = 0;

    private JPanel mainField;

    private JRadioButton radioButtonBlue;
    private JRadioButton radioButtonRed;
    private JRadioButton radioButtonYellow;
    private CanvasPanel canvasPanel1;


    public PictureForm() {

        canvasPanel1.addMouseListener(new MouseAdapter() {
            @Override
            public void mousePressed(MouseEvent e) {
                super.mousePressed(e);

                canvasPanel1.x1.add(e.getX());
                canvasPanel1.y1.add(e.getY());
            }

            @Override
            public void mouseReleased(MouseEvent e) {
                super.mouseReleased(e);

                canvasPanel1.x2.add(e.getX());
                canvasPanel1.y2.add(e.getY());

                if (radioButtonBlue.isSelected()) {
                    canvasPanel1.fuckYou.add(0);
                } else if (radioButtonRed.isSelected()) {
                    canvasPanel1.fuckYou.add(1);
                } else {
                    canvasPanel1.fuckYou.add(2);
                }

                length = canvasPanel1.x2.size();

                canvasPanel1.repaint();
            }
        });
    }

    public static void main(String[] args) {
        JFrame frame = new JFrame("PictureForm");
        frame.setContentPane(new PictureForm().mainField);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);
    }
}
