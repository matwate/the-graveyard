#include <iostream>
#include <fstream>
#include <string>

using namespace std;

string containsChar(string str, char c){
    for(int i = 0; i < str.size(); i++){
        if(str[i] == c){
            return str;
        }
    }
    return "";
}


int main(){

    ifstream file("letra.txt");

    if(!file.is_open()){
        cout << "No se pudo abrir el archivo" << endl;
        return 1;
    }

    string line;

    while(getline(file, line)){

        string word;

        for(int i = 0; i < line.size(); i++){
            if(line[i] == ' '){
                cout << containsChar(word, 'a') << endl;
                word = "";
            }else{
                word += line[i];
            }
        }

        
    }



    return 0;
}