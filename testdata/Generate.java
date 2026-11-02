import hqx.Hqx_4x;
import hqx.RgbYuv;
import java.io.*;
import java.util.Random;
import java.util.zip.GZIPOutputStream;

public class Generate {
 static DataOutputStream out;
 static void sample(int width, int height, boolean wx, boolean wy, int ty, int tu, int tv, int ta, int[] src) throws IOException {
  int[] dst = new int[width * height * 16];
  Hqx_4x.hq4x_32_rb(src, dst, width, height, ty, tu, tv, ta, wx, wy);
  out.writeInt(width); out.writeInt(height); out.writeBoolean(wx); out.writeBoolean(wy); out.writeInt(ty); out.writeInt(tu); out.writeInt(tv); out.writeInt(ta);
  for (int c : src) out.writeInt(c);
  for (int c : dst) out.writeInt(c);
 }
 public static void main(String[] args) throws Exception {
  RgbYuv.hqxInit();
  out = new DataOutputStream(new GZIPOutputStream(new FileOutputStream(args[0])));
  out.writeInt(768);
  for(int pattern = 0; pattern < 256; pattern++) {
   int[] src = new int[9]; int bit = 0;
   for(int i = 0; i < 9; i++) {
    if(i == 4) src[i] = 0xff123456;
    else src[i] = (pattern & (1 << bit++)) != 0 ? 0xfff0d0a0 : 0xff123456;
   }
   sample(3, 3, false, false, 48, 7, 6, 0, src);
  }
  Random random = new Random(73429);
  for(int n = 0; n < 512; n++) {
   int width=1+random.nextInt(5), height=1+random.nextInt(5);
   int alphaBase=(n & 4) != 0 ? 128 : 0;
   int[] palette=new int[5];
   for(int i=0;i<palette.length;i++) palette[i]=((alphaBase+random.nextInt(128))<<24)|random.nextInt(1<<24);
   int[] src=new int[width*height];
   for(int i=0;i<src.length;i++) src[i]=palette[random.nextInt(palette.length)];
   sample(width,height,(n&1)!=0,(n&2)!=0, n%3==0 ? 0 : 48, n%3==0 ? 0 : 7, n%3==0 ? 0 : 6, n%5==0 ? 64 : 0, src);
  }
  out.close();
 }
}
