#ifndef DIALOG_H
#define DIALOG_H

#include <QDialog>

class Dialog : public QDialog
{
    Q_OBJECT

public:
    Dialog(QWidget *parent = nullptr);
    ~Dialog();

    void ListWidgetDemo();
    void TreeWidgetDemo();
    void TableWidgetDemo();

};

#endif // DIALOG_H