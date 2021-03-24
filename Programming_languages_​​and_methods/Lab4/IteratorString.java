import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;

class Substring implements Iterable<String> {
    private ArrayList<String> strings = new ArrayList<>();

    // Конструктор
    Substring(String[] str) {
        strings.addAll(Arrays.asList(str));
    }
    //

    private class SubstringIterator implements Iterator<String> {
        private int point;

        SubstringIterator() {
            point = 0;
        }

        @Override
        public boolean hasNext() {
            if (strings.size() < 2) {
                throw new IllegalArgumentException("нужно 2 и более");
            }

            if (point < strings.size() - 1) {

                while (!((point < strings.size() - 1) && (strings.get(point + 1).lastIndexOf(strings.get(point)) >= 0))) {
                    if (strings.get(point).equals("")) {
                        point++;
                        System.out.println("Тут была путсота я ее пропустил :)");
                    }
                    if (strings.get(point + 1).equals("")) {
                        point++;
                        System.out.println("Тут была путсота я ее пропустил :)");
                    }

                    point++;
                }

                return (point < strings.size() - 1) && (strings.get(point + 1).lastIndexOf(strings.get(point)) >= 0);
            }
            else return false;
        }

        @Override
        public String next() {
            if (this.hasNext()) {
                return strings.get(++point);
            }
            return "That's all. The end";
        }
    }

    @Override
    public Iterator<String> iterator() {
        return new SubstringIterator();
    }
}

public class IteratorString {
    public static void main(String[] args) {

        String[] firstStr = {"abab", "ababab", "abababoba", "", "bob", "bobr"};
        Substring strings1 = new Substring(firstStr);

        String[] second ={"no", "non"};
        Substring strings2 = new Substring(second);

        System.out.println("Тестируем for");
        for (String s : strings1) {
            System.out.println(s);
        }

        System.out.println("Тестируем next()");
        Iterator<String> iter = strings2.iterator();

        System.out.println(iter.next());
        System.out.println(iter.next());
    }
}
