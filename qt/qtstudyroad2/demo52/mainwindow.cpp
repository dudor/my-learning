#include "mainwindow.h"
#include "QDragEnterEvent"
#include "qmimedata.h"
#include "QDropEvent"
#include "QList"
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
{
    this->textEdit = new QTextEdit();
    this->setCentralWidget(this->textEdit);

    this->textEdit->setAcceptDrops(false);
    this->setAcceptDrops(true);
    this->setWindowTitle("拖拽测试");

}

MainWindow::~MainWindow()
{
}

void MainWindow::dragEnterEvent(QDragEnterEvent *event)
{
    if(event->mimeData()->hasFormat("text/uri-list"))
    {
        event->acceptProposedAction();
    }
}

void MainWindow::dropEvent(QDropEvent *event)
{
    QList<QUrl> urls = event->mimeData()->urls();
    QString filename = urls.first().toLocalFile();
    if(filename.isEmpty())
        return;
    if(readFile(filename))
    {
        this->setWindowTitle(QString("%1 %2").arg(filename,"拖拽文件"));
    }

}

bool MainWindow::readFile(const QString &filename)
{
    QFile file(filename);
    QString content;
    if(file.open(QIODevice::ReadOnly))
    {
        content = file.readAll();
    }
    this->textEdit->setText(content);
}

