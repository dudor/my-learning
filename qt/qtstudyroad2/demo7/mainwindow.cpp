#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "qtoolbar.h"
#include "qmessagebox.h"
#include "qinputdialog.h"
#include "qdebug.h"
#include "qtextedit.h"
#include "qfiledialog.h"
#include "qfile.h"
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

       this->saveAction = new QAction(QIcon(":/images/bubble2"),"&Save",this);
        connect(this->saveAction,&QAction::triggered,this,&MainWindow::save);

       QMenu *menuFile = this->menuBar()->addMenu("&File");
       menuFile->addAction(this->openAction);

       QMenu *mSave = this->menuBar()->addMenu("&Save");
       mSave->addAction(saveAction);

       QToolBar *toolbar = this->addToolBar("File");
       toolbar->addAction(this->openAction);
       toolbar->addAction(saveAction);

       this->statusBar();
       this->textEdit = new QTextEdit(this);
       this->setCentralWidget(this->textEdit);

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
//    QInputDialog dialog(this);
//    dialog.exec();
//    qDebug() << dialog.textValue();

    QString path = QFileDialog::getOpenFileName(this,"Open File",".","*.txt");
    if(!path.isEmpty())
    {
        QFile file(path);
        if(!file.open(QIODevice::ReadOnly | QIODevice::Text))
        {
            QMessageBox::warning(this,"Read File","Cannt open file");
            return;
        }
        QTextStream in(&file);
        this->textEdit->setText(in.readAll());
        file.close();
    }
    else{
        QMessageBox::warning(this,"Read File","please select some files");
    }

}

void MainWindow::save()
{
    QString path = QFileDialog::getSaveFileName(this,"Save File",".","*.txt");
    if(!path.isEmpty()){
        QFile file(path);
        if(!file.open(QIODevice::WriteOnly | QIODevice::Text)){
            QMessageBox::warning(this,"Write File","Cannt open file");
            return;
        }
        QTextStream out(&file);
        out << this->textEdit->toPlainText();
        file.close();
    }
    else{
        QMessageBox::warning(this,"Write File","Please Select a file to save");
    }
}

MainWindow::~MainWindow()
{
    delete ui;
}

