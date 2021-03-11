class Card {
    private int value;
    private String suit;

    public Card(String suit, int value) {
        this.suit = suit;
        this.value = value;
    }

    public int getValue() {
        return value;
    }

    public String getSuit() {
        return suit;
    }

    public String toString() {
        return (value + " " + suit);
    }
}

class StackOfCards {
    private int maxSize;
    private String[] faces = {"Буби", "Крести", "Пики", "Черви"};
    private Card[] currentDeck;
    private int top;

    public StackOfCards(int maxSize) {
        this.maxSize = maxSize;
        currentDeck = new Card[maxSize];
        top = maxSize - 1;

        int tmpFace = 0;
        int currentValue = 6;
        for (int i = 0; i < maxSize; i++) {
            if (i % 4 == 0 && i != 0) {
                currentValue++;
            }
            currentDeck[i] = new Card(faces[tmpFace++ % 4], currentValue);
        }
    }

    public void shuffle() { // Алгоритм перетасовки Кнутта
        Card[] clonedDeck = currentDeck.clone();
        for(int i = 0; i < clonedDeck.length; i++)
        {
            int r = (int)(Math.random() * (clonedDeck.length - i) + i);

            Card tmp = clonedDeck[i];
            clonedDeck[i] = clonedDeck[r];
            clonedDeck[r] = tmp;
        }
        currentDeck = clonedDeck.clone();
    }

    public void showDeck() { // Вспомогательный метод для достоверности того, что все нормально заполняется)
        for (int i = 0; i < maxSize; i++) {
            if (currentDeck[i].getValue() == 11) {
                System.out.println("Валет " + " " + currentDeck[i].getSuit());
            } else if (currentDeck[i].getValue() == 12) {
                System.out.println("Дама " + " " + currentDeck[i].getSuit());
            } else if (currentDeck[i].getValue() == 13) {
                System.out.println("Король " + " " + currentDeck[i].getSuit());
            } else if (currentDeck[i].getValue() == 14) {
                System.out.println("Туз " + " " + currentDeck[i].getSuit());
            } else {
                System.out.println(currentDeck[i].getValue() + " " + currentDeck[i].getSuit());
            }
        }
    }

    public void push(Card card) {
        currentDeck[++top] = card;
    }

    public Card pop() {
        return currentDeck[top--];
    }

    public Card peek() {
        return currentDeck[top];
    }

}

class Checker {
    private Card[] combination;
    private boolean nothing = true;
    private boolean pair = false;
    private boolean three = false;
    private boolean street = false;
    private boolean flash = false;
    private boolean fullHouse = false;
    private boolean kare = false;
    private boolean flashRoyale = false;

    public Checker(Card[] combination) {
        this.combination = combination.clone();
        //System.out.println(this.combination[0].getValue());
    }

    public void searchForCombination() {

        int streetCount = 0;

        int[] allValueFinds = {0, 0, 0, 0, 0, 0, 0, 0, 0};
        int[] allSuitsFinds = {0, 0, 0, 0}; // БУБИ ЧЕРВИ КРЕСТИ ПИКИ

        for (int i = 0; i < 5; i++) {
            allValueFinds[combination[i].getValue() - 6]++;
            if (combination[i].getSuit().equals("Буби")) {
                allSuitsFinds[0]++;
            }
            if (combination[i].getSuit().equals("Черви")) {
                allSuitsFinds[1]++;
            }
            if (combination[i].getSuit().equals("Крести")) {
                allSuitsFinds[2]++;
            }
            if (combination[i].getSuit().equals("Пики")) {
                allSuitsFinds[3]++;
            }
        }

        for (int i : allValueFinds) { // для пары тройки и каре с фуллхаусом
            if (i == 2) {
                pair = true;
            }
            if (i == 3) {
                three = true;
            }
            if (pair == true && three == true) {
                fullHouse = true;
            }
            if (i == 4) {
                kare = true;
            }
        }

        for (int i : allSuitsFinds) { // для флэша
            if (i == 5) {
                flash = true;
            }
        }
        // проверки на флеш рояль нет, но надо просто проверять, выпало ли 5 старших карт при проверке стрита
        for (int i = 0; i + 5 < allValueFinds.length; i++) {
            for (int j = i; j < i + 5; j++) {
                if (allValueFinds[j] > 0) {
                    streetCount++;
                } else {
                    streetCount = 0;
                    break;
                }
            }
            if (streetCount >= 5) {
                street = true;
                break;
            }
        }

        if (flash == true && street == true) {
            flashRoyale = true;
        }
    }

    public void showCombo() {
        if (flashRoyale == true) {
            System.out.println("FullHouse");
        } else if (kare == true) {
            System.out.println("Kare");
        } else if (fullHouse == true) {
            System.out.println("FH");
        } else if (flash == true) {
            System.out.println("Flash");
        } else if (street == true) {
            System.out.println("Street");
        } else if (three == true) {
            System.out.println("Set");
        } else if (pair == true) {
            System.out.println("Pair");
        } else  {
            System.out.println("Poor");
        }
    }
}

public class Poker {
    public static void main(String[] args) {
        StackOfCards deck = new StackOfCards(36);
        deck.shuffle();
        Card[] cards = new Card[5];

        for (int i = 0; i < 5; i++) {
            cards[i] = deck.pop();
            System.out.println(cards[i].toString());
        }

        Checker checker = new Checker(cards);
        checker.searchForCombination();
        checker.showCombo();
    }
}
