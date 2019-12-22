#include <QCoreApplication>
#include "newspaper.h"
#include "reader.h"
int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);
    Newspaper np("newspaper 1");
    Reader reader;
    QObject::connect(&np,&Newspaper::newPaper,&reader,&Reader::receiveNewspaper);

    np.send();
    return a.exec();
}
