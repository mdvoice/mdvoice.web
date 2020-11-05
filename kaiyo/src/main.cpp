#include <iostream>
#include <fstream>

using namespace std;
int main(int argc, char const *argv[])
{
   // cout<<argc;
   if (argc == 1)
   {
      while (1)
      {
         int  i = 0;
         char a[10][100];
         cout << "kaiyo@cli >";
         while (cin >> a[i])
         {
            cout << a[i] << endl;
         }
         // for (size_t i = 0; i < 10; i++) {
         // cin>>a[i];
         // cout<<a[i]<<endl;
         // if (a[i][0]=='\n') {
         //   cout<<"end!"<<endl;
         //   break;
         // }
         // }
      }
   }


   return 0;
}
