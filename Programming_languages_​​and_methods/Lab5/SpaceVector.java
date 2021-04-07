import java.util.ArrayList;
import java.util.Arrays;
import java.util.Optional;
import java.util.stream.Stream;

public class SpaceVector {
    final ArrayList<SingleVector> v = new ArrayList<>();
    SingleVector tmp;

    public SpaceVector(int nel, int dimension) {
        SingleVector[] help = new SingleVector[nel];

        Arrays.stream(help).forEach(x -> {
            x = new SingleVector(dimension);
            v.add(x);
        });

        tmp = v.get(0);
    }


    public ArrayList<SingleVector> getV() {
        return v;
    }


    // возвращает поток нормализированных векторов
    public Stream<SingleVector> normalizeAllVectors() {
        return v.stream().map(SingleVector::normalizeVector);
    }

    // возвращает вектор, ортогональный всем остальным, если он не NULL
    public Optional<SingleVector> findOrthogonal() {
        v.forEach(vec1 -> { // находим усредненный скаляр
            vec1.findAverageScalar(this);
            //System.out.println("Scalar inside - " + vec1.getAverageScalar());
        });

        // вот он !!!
        v.forEach(x -> {
            if (Math.abs(x.getAverageScalar()) < Math.abs(tmp.getAverageScalar())) {
                tmp = x;
            }
        });

        return Optional.ofNullable(tmp);
    }

    // вспомогательный метод для вывода векторов на экран
    public void showSpace() {
        v.forEach(SingleVector::showVector);
    }
}
