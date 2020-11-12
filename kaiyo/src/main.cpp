#include <iostream>
#include <fstream>
#include <string>
// #include <yaml.h>
#include "include/yaml-cpp/yaml.h"


using namespace std;
int main(int argc, char const *argv[])
{
   // cout<<argc;
   if (argc == 1)
   {
      while (1)
      {
         int  i = 0;
         string a;
         cout << "kaiyo@cli >";
         getline(cin,a,'\n');
         if(a!=""){
           YAML::Node config = YAML::LoadFile(a);

           // cout<<config.files<<endl;
         }
         // while (cin.good())
         // {
         //    cin >> a[i];
         //    cout << a[i] << endl;
         // }
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
