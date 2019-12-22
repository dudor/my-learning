#ifndef READER_H
#define READER_H

#include <QObject>

class Reader : public QObject
{
    Q_OBJECT
public:
    explicit Reader(QObject *parent = nullptr);
    void receiveNewspaper(const QString &name) const;

signals:




};

#endif // READER_H
