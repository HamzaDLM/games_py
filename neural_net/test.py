a = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

rsize = 3
csize = 4

for r in range(rsize):
    for c in range(csize):
        print(a[r * csize + c], end=" ")
    print("")


# def print_matrix(array, rows, columns):
#     if len(array) != rows * columns:
#         print("Error: Invalid array size for the given number of rows and columns.")
#         return

#     for i in range(rows):
#         for j in range(columns):
#             index = i * columns + j
#             print(array[index], end="\t")
#         print()


# # Example usage
# array = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rows = 3
# columns = 4

# print_matrix(array, rows, columns)
