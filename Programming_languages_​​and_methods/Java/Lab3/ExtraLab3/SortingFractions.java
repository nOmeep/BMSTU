import java.util.Arrays;

class Fraction implements Comparable<Fraction>{
    final int numerator;
    final int denominator;
    final int currentGcd;

    public Fraction(int numerator, int denominator) {
        if (denominator == 0) {
            throw new IllegalArgumentException("Denominator cannot be zero");
        }
        currentGcd = gcd(numerator, denominator);
        this.numerator = numerator / currentGcd;
        this.denominator = denominator / currentGcd;
    }

    ///////////////////////////
    public int getNumerator() {
        return numerator;
    }

    public int getDenominator() {
        return denominator;
    }
    ////////////////////////////

    public int gcd(int a, int b) {
        if (b == 0) {
            return a;
        }
        return gcd(b, a % b);
    }

    @Override
    public int compareTo(Fraction fraction) {
        // local variables for calc
        int compareGcd = gcd(this.getDenominator(), fraction.getDenominator());
        int firstMul = this.getDenominator() / compareGcd;
        int secondMul = fraction.getDenominator() / compareGcd;

        return (this.getNumerator() * secondMul - fraction.getNumerator() * firstMul);
    }

    @Override
    public String toString() {
        return (getNumerator() + "/" + getDenominator());
    }
}

public class SortingFractions {
    public static void main(String[] args) {
        Fraction[] fractions = new Fraction[5];

        for (int i = 0; i < fractions.length; i++) {

            // local numerator and denominator
            int num = (int)(Math.random() * 9 + 1);
            int den = (int)(Math.random() * (num * 2 - 1) + num);

            fractions[i] = new Fraction(num, den);
            System.out.println(fractions[i] + " before sort");
        }
        System.out.println("Sorted");
        Arrays.sort(fractions);

        //output
        for (Fraction f: fractions) {
            System.out.println(f + " after sort");
        }
    }
}
