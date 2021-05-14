:h1:`Bootstrap RST Cheatsheet`

Bootstrap RST is extensive; as is the RST spec itself. This document serves 
as a curated cheatsheet for what I consider to be the most useful features.

Here are some more comprehensive references:

* `Comprehensive Bootstrap RST CSS Reference <http://rougier.github.io/bootstrap-rst/CSS.html>`_
* `Comprehensive Bootstrap RST Component Reference <http://rougier.github.io/bootstrap-rst/components.html>`_
* `General RST reference <http://docutils.sourceforge.net/docs/user/rst/quickref.html>`_

.. container:: bs-callout

   If you would like direct access to the rst source file that generated this cheatsheet,
   you can grab it from the container directly:
   
   ``docker copy <name-of-container>:/srv/rst-boilerplate/cheatsheet.rst /path/to/cheatsheet.rst``

:h1:`RST <small>with bootstrap</small>`

:h2:`Simple formatting`

**bold** *italic* ``literal`` :code:`code`
::

    **bold** *italic* ``literal`` :code:`code`
    ::
        
        <literal block>

Bullet list (nesting requires **empty lines**)

- apples
- bananas
  
  + grapes
  + cherries

- avacodos

::

    - apples
    - bananas
  
       - grapes
       - cherries

    - avacodos

Numbered list

1. first
#. second
#. third

::

   1. first
   #. second
   #. third

:h2:`References`

RST is great for cross referencing!

:h3:`External <small>(different site)</small>`

`Bootstrap RST <http://rougier.github.io/bootstrap-rst/CSS.html#helper-classes>`_

::

    `Bootstrap RST <http://rougier.github.io/bootstrap-rst/CSS.html#helper-classes>`_

:h3:`External <small>(current site)</small>`

Internal `sample <sample.html>`_

::

    Internal `sample <sample.html>`_

:h3:`Internal <small>(current site)</small>`

Here is an internal _`anchor`

Goes to the `anchor`_

::

    Here is an internal _`anchor`

    Goes to the `anchor`_

.. _`invisible anchor`:

And this is an invisible anchor which comes from the directive above

This goes to the `invisible anchor`_

::

   .. _`invisible anchor`:

   And this is an invisible anchor which comes from the directive above

   This goes to the `invisible anchor`_
   
:h2:`Comments`

Add a comment like this:

::

    .. comment comment comment
       comment comment

:h2:`Blocks`

.. container:: bs-callout

   foo bar ``literal`` :code:`code`

.. callout:: warning

   foo bar ``literal`` :code:`code` 

.. admonition:: :h3:`Alternate elements`
   :class: bs-callout bs-callout-info

   foo bar :code:`<div>some code</div>`

.. container:: bs-example

   foo bar

.. code::

   .. container:: bs-callout

      foo bar ``literal`` :code:`code`

   .. callout:: warning

      foo bar ``literal`` :code:`code`
    
   .. NOTE: on above... instead of "warning", also try "danger" and "info"

   .. container:: bs-example

      foo bar

   .. admonition:: :h4:`Alternate elements`
      :class: bs-callout bs-callout-info

      foo bar :code:`<div>some code</div>`

:h2:`Headings`

.. container:: bs-example

    :h1:`h1. Bootstrap heading <small>sub heading</small>`
    :h2:`h2. Bootstrap heading`
    :h3:`h3. Bootstrap heading`
    :h4:`h4. Bootstrap heading`
    :h5:`h5. Bootstrap heading`
    :h6:`h6. Bootstrap heading`

::

    :h1:`h1. Bootstrap heading <small>sub heading</small>`
    :h2:`h2. Bootstrap heading`
    :h3:`h3. Bootstrap heading`
    :h4:`h4. Bootstrap heading`
    :h5:`h5. Bootstrap heading`
    :h6:`h6. Bootstrap heading`


.. callout:: warning

  Avoid using the headings with ---- or ===== or whatever.
  They seem to cause issues sometimes.

:h2:`Tables`

**List Table**

.. container::

   .. list-table::
      :widths: 75 25
      :class: table

      * - :h4:`h4. Bootstrap heading`
        - **Semibold 18 px**
      * - :h5:`h5. Bootstrap heading`
        - **Semibold 14 px**
      * - :h6:`h6. Bootstrap heading`
        - **Semibold 12 px**

::

    .. container::

       .. list-table::
          :widths: 75 25
          :class: table

          * - :h4:`h4. Bootstrap heading`
            - **Semibold 18 px**
          * - :h5:`h5. Bootstrap heading`
            - **Semibold 14 px**
          * - :h6:`h6. Bootstrap heading`
            - **Semibold 12 px**

+------------+------------------+--------------+
| **Ticket** | **Description**  | **Assignee** |
+------------+------------------+--------------+
| ABC-123    | Easy as do re mi | Me           |
+------------+------------------+--------------+

::

    +------------+------------------+--------------+
    | **Ticket** | **Description**  | **Assignee** |
    +------------+------------------+--------------+
    | ABC-123    | Easy as do re mi | Me           |
    +------------+------------------+--------------+


:h1:`Bootstrap RST Components`

:h2:`Labels`

.. container:: bs-example

   :label-default:`Default`
   :label-primary:`Primary`
   :label-success:`Success`
   :label-info:`Info`
   :label-warning:`Warning`
   :label-danger:`Danger`

.. code::
   :class: highlight

   :label-default:`Default`
   :label-primary:`Primary`
   :label-success:`Success`
   :label-info:`Info`
   :label-warning:`Warning`
   :label-danger:`Danger`


:h2:`Breadcrumbs`

.. class:: breadcrumb

   * Home
   * Library
   * Data

::
    
   .. class:: breadcrumb

      + Home
      + Library
      + Data
