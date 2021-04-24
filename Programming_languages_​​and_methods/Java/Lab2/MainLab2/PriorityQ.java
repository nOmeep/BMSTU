class PriorityQueue {
    private int maxSize;
    private long[] queArray;
    private int nItems;

    public PriorityQueue(int s) {
        maxSize = s;
        queArray = new long[maxSize];
        nItems = 0;
    }

    public void insert(long item) {
        int j;

        if(nItems == 0) { // если пусто
            queArray[nItems++] = item;
        } else { // если не пусто :)
            for (j = nItems - 1; j >= 0; j--) { // перебираем в обратном направлении
                if (item > queArray[j]) {
                    queArray[j + 1] = queArray[j];
                } else {
                    break;
                }
            }
            queArray[j + 1] = item; //вставка
            nItems++; // увеличиваем
        }
    }

    public long remove() {
        return queArray[--nItems];
    }

    public int countNels() {
        return nItems;
    }

    public long peekMax() {
        return queArray[0];
    }

    public boolean empty() {
        return (nItems == 0);
    }

    public boolean isFull() {
        return (nItems == maxSize);
    }
}

public class PriorityQ {
    public static void main(String[] args) {
        PriorityQueue pq = new PriorityQueue(5); // fixed size

        pq.insert(30);
        System.out.println("Колчество в очереди - " + pq.countNels() + "\n");
        pq.insert(50);
        pq.insert(10);
        pq.insert(40);
        System.out.println("Колчество в очереди - " + pq.countNels() + "\n");
        pq.insert(20);

        System.out.println(pq.peekMax() + "\n");

        while (!pq.empty()) {
            long item = pq.remove();
            System.out.println(item + " ");
        }
        System.out.println("");
    }
}
