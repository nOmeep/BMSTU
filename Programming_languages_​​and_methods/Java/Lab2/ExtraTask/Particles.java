class Universe {
    private static int pCount = (int)(1000 + Math.random() * 10000);
    
    public static int getCount() { // выдает количество частиц во вселенной :)
        return pCount;
    }
}

class Molecules extends Universe {
    // скорость света известна частицам и неизменна
    private static double[] mWeight;
    private static double[] singleKinetic;
    
    public Molecules() {
        mWeight = new double[getCount()];
        singleKinetic = new double[getCount()];
        
        for (int i = 0; i < getCount(); i++) { // каждой частице свое случайное значение массы
            mWeight[i] = Math.random() * 2;
        }
    }

    public double[] allKineticPowers() {
        for (int i = 0; i < getCount(); i++) {
            double pSpeed = 1 + Math.random() * 3;
            singleKinetic[i] = mWeight[i] * (pSpeed * pSpeed) / 2;
        }
        return singleKinetic;
    }
}

public class Particles {
    public static void main(String[] args) {
        double result = 0;
        Molecules mol = new Molecules();
        for (int i = 0; i < mol.getCount(); i++) {
            result += mol.allKineticPowers()[i];
        }
        System.out.println(result / mol.getCount());
    }
}



