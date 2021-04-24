import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class SingleVector {
    private List<Double> sv;
    private double lengthVector;

    public SingleVector(int dimension) {
        double[] coordinates = new double[dimension];
        sv = new ArrayList<>();
        Arrays.stream(coordinates).forEach(x -> {
            x = Math.random() * 20 - 10;
            sv.add(x);
        });

        // значальная длина вектора для вывода
        sv.forEach(x -> lengthVector += x * x);
        lengthVector = Math.sqrt(lengthVector);
    }

    public SingleVector normalizeVector() {
        sv = sv.stream().map(x -> x / lengthVector).collect(Collectors.toList());

        // я честно пересчитываю длину при нормализации
        lengthVector = 0;
        sv.forEach(x -> lengthVector += x * x);
        lengthVector = Math.sqrt(lengthVector);

        return this;
    }

    public List<Double> getSv() {
        return sv;
    }

    // считаем среднее скалярное произведение
    private double averageScalar = 0;
    public void findAverageScalar(SpaceVector space) {
        space.getV().stream().filter(x -> x != this).forEach(vec1 -> vec1.getSv().forEach(c -> this.getSv().forEach(z -> averageScalar += (c * z))));
        averageScalar /= space.getV().size();
    }

    public double getAverageScalar() {
        return averageScalar;
    }

    // Вспомогательный метод для вывода координат вектора
    public void showVector() {
        sv.forEach(x -> System.out.printf("Координата %2.2f\n", x));
        System.out.printf("Длина - %2.2f\n", lengthVector);
    }
}
