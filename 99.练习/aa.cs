using System;
using System.Collections;
using System.Collections.Generic;
namespace test
{
    class Class1
    {
        static void Main(string[] args)
        {
            
            string str = Console.ReadLine();
            while (!String.IsNullOrEmpty(str))
            {
                int num = Convert.ToInt32(str);
                List<long> q2 = new List<long>();
                List<long> q3 = new List<long>();
                List<long> q5 = new List<long>();
                //初始化
                q2.Add(2); q3.Add(3); q5.Add(5);
                long temp = 0;
                for (int i = 1; i < num; i++)
                {
                    temp = getQMin(q2, q3, q5);
                }
                Console.WriteLine(temp);
                str = Console.ReadLine();
            }
            
        }
        public static long getQMin(List<long> q2, List<long> q3, List<long> q5)
        {
            long qmin = Convert.ToInt64(q2[0]);
            int flag = 2;
            if(qmin > Convert.ToInt64(q3[0]))
            {
                qmin = Convert.ToInt64(q3[0]);
                flag = 3;
            }
            if (qmin > Convert.ToInt64(q5[0]))
            {
                qmin = Convert.ToInt64(q5[0]);
                flag = 5;
            }
            if (flag == 2)
                q2.Remove(qmin);
            if (flag == 3)
                q3.Remove(qmin);
            if (flag == 5)
                q5.Remove(qmin);
            //更新数据	
            q5.Add(5 * qmin);
            if (flag < 5)
                q3.Add(3 * qmin);
            if (flag < 3)
                q2.Add(2 * qmin);
            return qmin;
        }
    }
    
}