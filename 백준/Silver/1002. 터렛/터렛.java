import java.util.*;

public class Main {

	public static void main(String[] args) {
		Scanner sc=new Scanner(System.in);
		int N=sc.nextInt();
		
		for(int i=0; i<N;i++) {
			int x1=sc.nextInt();
			int y1=sc.nextInt();
			int r1=sc.nextInt();
			
			int x2=sc.nextInt();
			int y2=sc.nextInt();
			int r2=sc.nextInt();
			
			double D=Math.sqrt((x1-x2)*(x1-x2)+(y1-y2)*(y1-y2));
			if(D==0&&r1==r2) System.out.println(-1);
			else if(Math.abs(r1-r2)>D||D>Math.abs(r1+r2)) System.out.println(0);
			else if(D==Math.abs(r1+r2)) System.out.println(1);
			else if(D==Math.abs(r1-r2)) System.out.println(1);
			else if(D<Math.abs(r1+r2)) System.out.println(2);
		}
	}
}
