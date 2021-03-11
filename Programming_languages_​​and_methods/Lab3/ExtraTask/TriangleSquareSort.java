import java.util.Arrays;

class Point {
    private final double x;
    private final double y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public double getX() {
        return x;
    }

    public double getY() {
        return y;
    }
}

class Triangle implements Comparable<Triangle> {
    private final int square;

    public Triangle() {
        Point[] _points = new Point[3];

        for (int i = 0; i < 3; i++) {
            _points[i] = new Point((int)(Math.random() * 9 + 1), (int)(Math.random() * 9 + 1));
        }

        double ab = Math.sqrt(Math.pow(_points[1].getX() - _points[0].getX(), 2) + Math.pow(_points[1].getY() - _points[0].getY(),2));
        double bc = Math.sqrt(Math.pow(_points[2].getX() - _points[1].getX(), 2) + Math.pow(_points[2].getY() - _points[1].getY(),2));
        double ac = Math.sqrt(Math.pow(_points[2].getX() - _points[0].getX(), 2) + Math.pow(_points[2].getY() - _points[0].getY(),2));
        double p = (ab + bc + ac) / 2;

        square = (int)(Math.sqrt(p * (p - ac) * (p - bc) * (p - ab)));
    }

    public int getSquare() {
        return square;
    }

    @Override
    public String toString() {
        return ("Площадь треугольника равна - " + square);
    }

    @Override
    public int compareTo(Triangle triangle) {
        return (this.getSquare() - triangle.getSquare());
    }
}

public class TriangleSquareSort {
    public static void main(String[] args) {
        Triangle[] triangles = new Triangle[3];
        for (int i = 0; i < 3; i++) {
            triangles[i] = new Triangle();
            System.out.println(triangles[i]);
        }
        System.out.println("После сортировки ");
        Arrays.sort(triangles);
        for (int i = 0; i < 3; i++) {
            System.out.println(triangles[i]);
        }
    }
}
