public class Test {
    public static void main(String[] args) {
        SpaceVector space = new SpaceVector(4, 3);

        System.out.println("Before normalize \n");
        space.showSpace();

        System.out.println("\nAfter normalize \n");
        space.normalizeAllVectors()./*filter(x -> Double.compare(x.getLengthVector(), 1.0f) == 0).*/forEach(SingleVector::showVector); // фильтр позволяет выводить только единичные вектора

        System.out.println("\nVector orthogonal to the rest\n");
        space.findOrthogonal().ifPresent(x -> x.showVector());

    }
}