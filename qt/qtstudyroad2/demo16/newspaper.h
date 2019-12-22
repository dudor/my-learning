#ifndef NEWSPAPER_H
#define NEWSPAPER_H

#include <QObject>

class Newspaper : public QObject
{
    Q_OBJECT
public:
    explicit Newspaper(QObject *parent = nullptr);
    explicit Newspaper(const QString &name);

    void send() const;
signals:
    void newPaper(const QString &name) const;

private:
    QString m_name;

};

#endif // NEWSPAPER_H
