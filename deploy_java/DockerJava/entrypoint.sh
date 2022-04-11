#java -jar $JAVA_OPT  /usr/local/erp-ccis.jar $START_OPT > /logs/ccis.out
java -jar $JAVA_OPT  /usr/local/app.jar $START_OPT | tee /usr/local/app.out ; exit ${PIPESTATUS[0]}