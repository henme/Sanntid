# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread

i = 0

def someThreadFunction():
	global i
	for j in range(0,1000000):
	    i += 1;

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
    

def someThreadFunction2():
    global i
    for j in range(0,1000000):
    	i -= 1;

def main():
    someThread = Thread(target = someThreadFunction, args = (),)
    someThread2 = Thread(target = someThreadFunction2, args = (),)
    someThread.start()
    someThread2.start()
    someThread.join()
    someThread2.join()
    print(i)

main()