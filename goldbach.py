""" This module demonstrate the Goldbach conjecture. """

import sys

def get_primes(max_number):
    """ Computes and returns all prime numbers up to max_number. """
    primes = []
    for value in range(2, max_number + 1):
        is_prime = True
        for prime in primes:
            if value % prime == 0:
                is_prime = False
                break
        if is_prime:
            primes.append(value)
    return primes


def goldbach(value):
    """ Returns a list of integer such that for each p in the list,
        p + q = value, p <= q, and both p and q are primes. """
    result = []
    if value >= 4 and value % 2 == 0:
        primes = get_primes(value)
        for prime in primes:
            if prime > value / 2:
                break
            difference = value - prime
            if difference in primes:
                result.append(prime)
    return result


def main():
    """ Entry point of the program. """
    data = []
    if len(sys.argv) > 1:
        for filename in sys.argv[1:]:
            data.extend(readfile(filename))
    else:
        data = (3, 4, 14, 26, 100)
    for value in data:
        print_goldbach(value, goldbach(value))


def print_goldbach(value, prime_list):
    """ Given the value of n and the output of goldbach(value),
        prints out the result to the screen. """
    if not prime_list:
        print('We found no Goldbach pairs for %d.' % value)
    else:
        print('We found %d Goldbach pair(s) for %d.' %
              (len(prime_list), value))
        for prime in prime_list:
            difference = value - prime
            print('%d = %d + %d' % (value, prime, difference))
    print('')


def readfile(filename):
    """ Reads and returns data from a file,
        we assume no problem on opening and reading from the file. """
    data = []
    file_handle = open(filename)
    for line in file_handle:
        data.append(int(line))
    file_handle.close()
    return data


main()  # Invoke the main program.
