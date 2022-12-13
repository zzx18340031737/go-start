package com.test;

public class Ref {
public static void main(String[] args) {
//1.通过对象调用getClass()方法来获取，通常应用在比如你传过来一个Object类型的对象，而我不知道你具体是什么类的情况
Student stul = new Studnet();
Class stuClass = stul.getClass();
//2.直接通过"类名.Class"的方式得到，该方法最为安全可靠，程序性能更高
//这说明任何一个类都有一个隐含的静态成员变量class
Class stuClass = Student.class
//3.通过class对象的forName（）静态方法来获取，用得最多，但可能抛出ClassNotFound Exception异常
try {
  //注意此字符串必须是真实路径，就是带包名的类路径"包名.类名"
  Class stuClass3 = Class.forName("com.test.Studnet")
} catch (ClassNotFoundException e) {
e.PrintStackTrace();
}
}
}
