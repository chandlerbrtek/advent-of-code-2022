package day7;

import java.io.BufferedReader;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class DaySeven {
    public static void main(String[] args) throws IOException {

        FileInputStream fstream = new FileInputStream("day7/7.txt");
        BufferedReader br = new BufferedReader(new InputStreamReader(fstream));

        String strLine;
        HashMap<String, Integer> cache = new HashMap<>();
        List<String> path = new ArrayList<>();

        //Read File Line By Line
        while ((strLine = br.readLine()) != null)   {
            String[] tokens = strLine.split(" ");

            if (tokens[0].equals("$")) {
                if (tokens[2].equals("..")) {
                    path.remove(path.size() - 1);
                } else {
                    path.add(tokens[2]);

                    StringBuilder sb = new StringBuilder();
                    for (String p : path) {
                        sb.append(p);
                        sb.append("/");
                    }

                    if (!cache.containsKey(sb.toString())) {
                        cache.put(sb.toString(), 0);
                    }
                }
            } else {
                StringBuilder sb = new StringBuilder();
                for (String p : path) {
                    sb.append(p);
                    sb.append("/");

                    cache.put(sb.toString(), cache.get(sb.toString()) + Integer.parseInt(tokens[0]));
                }
            }

        }

        int total = 0;

        for (int val : cache.values()) {
            if (val <= 100000) {
                total += val;
            }
            if (val >= 8748071) {
                System.out.println(val);
            }
        }

        System.out.println(total);


        //Close the input stream
        fstream.close();

    }
}