#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "qtoolbar.h"
#include "qmessagebox.h"
#include "qinputdialog.h"
#include "qdebug.h"
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
       this->setWindowTitle("MainWindow");
       this->openAction = new QAction(QIcon(":/images/bubble2") , "&Open File",this);
       this->openAction->setShortcut(QKeySequence::Open);
       this->openAction->setStatusTip("Open an existing file.");
       this->connect(this->openAction,&QAction::triggered,this,[=]{this->open();});
       QMenu *menuFile = this->menuBar()->addMenu("&File");
       menuFile->addAction(this->openAction);

       QToolBar *toolbar = this->addToolBar("File");
       toolbar->addAction(this->openAction);

       this->statusBar();

}

void MainWindow::open()
{
    //QMessageBox::information(this,"提示","打开文件");
//    QDialog dialog(this);
//    dialog.setWindowTitle("This is a dialog");
//    dialog.exec();

//    QDialog *dialog = new QDialog(this);
//    dialog->setAttribute(Qt::WA_DeleteOnClose);
//    dialog->setWindowTitle("This is a dialog");
//    dialog->exec();
    QInputDialog dialog(this);
    dialog.exec();
    qDebug() << dialog.textValue();
}

MainWindow::~MainWindow()
{
    delete ui;
}

