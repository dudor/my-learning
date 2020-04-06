#ifndef DIALOG_H
#define DIALOG_H

#include <QDialog>
#include "QItemSelection"
#include "qtablewidget.h"
QT_BEGIN_NAMESPACE
namespace Ui { class Dialog; }
QT_END_NAMESPACE

class Dialog : public QDialog
{
    Q_OBJECT

public:
    Dialog(QWidget *parent = nullptr);
    ~Dialog();

    QTableWidget* tb;
    void updateSelection(const QItemSelection &selected, const QItemSelection &deselected);
    void changeCurrent(const QModelIndex &current,const QModelIndex &previous);


private:
    Ui::Dialog *ui;
};
#endif // DIALOG_H
