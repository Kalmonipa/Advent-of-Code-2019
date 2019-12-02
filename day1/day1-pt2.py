
def calculateMass(mass,total):
    
    if mass <=0 :
        return total
    divideThree = mass / 3
    if divideThree <= 0 :
        return total
    minusTwo = (divideThree - 2)
    if minusTwo <= 0 :
        return total
    
    total = total + minusTwo

    if minusTwo > 6 :
        calculateMass(minusTwo, total)
    else:
        return total

    # print 'Line:' + str(line)
    # print 'Divided by 3: ' + str(divideThree)
    # print 'Minus 2:' + str(minusTwo)


def main():

    filepath = 'input'
    inputFile = open(filepath)

    total = 0

    for line in inputFile.readlines():

        num = int(line)

        total = calculateMass(num,total)


        # line = int(line)
        # divideThree = line / 3
        # minusTwo = (divideThree - 2)
        
        #total = total + minusTwo

        #print 'Line:' + str(line)
        #print 'Divided by 3: ' + str(divideThree)
        #print 'Minus 2:' + str(minusTwo)

    print total
main()


