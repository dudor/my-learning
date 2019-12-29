#include <QCoreApplication>
#include "QDate"
#include "qlist.h"
#include "qdatastream.h"
#include "QListIterator"
#include "qdebug.h"
#include "QMutableListIterator"
struct Movie{
    int id;
    QString title;
    QDate releaseDate;
};

QDataStream &operator<<(QDataStream &out,const Movie &m){
     out << (quint32)m.id << m.title << m.releaseDate;
     return out;
}

QDataStream &operator>>(QDataStream &in,Movie &m){
    quint32 id;
    QDate date;
    in >> id >> m.title >> date;
    m.id = id;
    m.releaseDate = date;
    return in;

}



int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);
    QList<Movie> movs;
    QList<QString> list;
    list<<"a"<<"b"<<"c"<<"d";
    QListIterator<QString> i(list);
    while (i.hasNext()) {
        qDebug() << i.next();
    }

    QMutableListIterator<QString> m(list);
    while (m.hasNext()) {
        QString v = m.next();
        v = v.toUpper();
        qDebug() << v;
    }

    QList<QString>::iterator ii;
    for(ii = list.begin();ii != list.end();ii ++){
        *ii = (*ii).toUpper() + (*ii).toUpper();
        qDebug()<< *ii;
    }

    foreach(const QString &str,list){
        qDebug()<< str;
    }



    return a.exec();
}



